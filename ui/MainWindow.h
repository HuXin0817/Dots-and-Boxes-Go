#pragma once

#include <QApplication>
#include <QColor>
#include <QTimer>
#include <chrono>
#include <thread>

#include "../core/ai/AIConfig.h"
#include "BoxCanvasLayer.h"
#include "DotCanvasLayer.h"
#include "EdgeButtonLayer.h"
#include "EdgeCanvasLayer.h"

class MainWindow final : public BaseCanvasLayer {
  Q_OBJECT
  public:
  MainWindow(bool AIPlayer1,
             bool AIPlayer2,
             AIModelType AIPlayer1Type,
             AIModelType AIPlayer2Type,
             QWidget* parent = nullptr)
      : BaseCanvasLayer(parent),
        AIPlayer1(AIPlayer1),
        AIPlayer2(AIPlayer2),
        Player1Model(AIConfig::createModel(AIPlayer1Type)),
        Player2Model(AIConfig::createModel(AIPlayer2Type)) {
    resize(WindowSize, WindowSize);
    setMinimumSize(WindowSize, WindowSize);

    Board.New();
    BoxCanvasLayer.New(this);
    EdgeCanvasLayer.New(this);
    DotCanvasLayer.New(this);
    std::function<std::function<void()>(Edge)> CallBackFactory = [this](Edge edge) {
      return [edge, this] { return setPlayerMoveEdge(edge); };
    };
    EdgeButtonLayer.New(CallBackFactory, this);
  }

  QColor
  Color() const override {
    if (isDarkMode()) {
      return {43, 43, 43, 255};
    } else {
      return {242, 242, 242, 255};
    }
  }

  signals:
  void
  requestAIMove();

  public slots:
  void
  Add(Edge edge) {
    if (Board->NowStep() > 0) {
      EdgeCanvasLayer->Canvases.At(LastEdge)->highLight = false;
    }
    EdgeCanvasLayer->Canvases.At(edge)->state = StateFromTurn(Board->Turn);
    EdgeCanvasLayer->Canvases.At(edge)->raise();

    for (Box box : EdgeBoxMapper::EdgeNearBoxes.At(edge)) {
      int count = 0;
      for (Edge nearEdge : EdgeBoxMapper::BoxNearEdges.At(box)) {
        if (Board->Contains(nearEdge)) {
          count++;
        }
      }
      if (count == 3) {
        BoxCanvasLayer->BoxCanvases.At(box)->state = StateFromTurn(Board->Turn);
      }
    }

    Board->Add(edge);
    LastEdge = edge;
    update();
    QApplication::beep();
  }

  protected:
  void
  paintEvent(QPaintEvent* event) override {
    BaseCanvasLayer::paintEvent(event);

    QPainter painter(this);
    painter.fillRect(rect(), Color());
  }

  void
  resizeEvent(QResizeEvent* event) override {
    BaseCanvasLayer::resizeEvent(event);

    int x = (width() - WindowSize) / 2;
    int y = (height() - WindowSize) / 2;

    BoxCanvasLayer->move(x, y);
    EdgeCanvasLayer->move(x, y);
    DotCanvasLayer->move(x, y);
  }

  void
  showEvent(QShowEvent* event) override {
    BaseCanvasLayer::showEvent(event);

    std::thread([this] {
      while (Board->Gaming()) {
        if (AIPlayer1 && Board->Turn == Player1Turn) {
          PlayerMoveEdge = RandomChoice(Player1Model->BestCandidateEdges(*Board));
        } else if (AIPlayer2 && Board->Turn == Player2Turn) {
          PlayerMoveEdge = RandomChoice(Player2Model->BestCandidateEdges(*Board));
        } else {
          PlayerMoveEdge = -1;
          while (PlayerMoveEdge == -1) {
            std::this_thread::yield();
          }
        }
        Add(PlayerMoveEdge);

        printf("| Step %d | Player %d Move (%d, %d) -> (%d, %d) | Score %d : %d |\n",
               Board->NowStep(),
               Board->Turn == Player1Turn ? 1 : 2,
               PlayerMoveEdge.Dot1().X(),
               PlayerMoveEdge.Dot1().Y(),
               PlayerMoveEdge.Dot2().X(),
               PlayerMoveEdge.Dot2().Y(),
               Board->Player1Score,
               Board->Player2Score);
      }

      if (Board->Player1Score > Board->Player2Score) {
        printf("| Player 1 Win! |\n");
      } else if (Board->Player2Score > Board->Player1Score) {
        printf("| Player 2 Win! |\n");
      } else {
        printf("| Draw! |\n");
      }

      std::this_thread::sleep_for(std::chrono::seconds(2));
      EdgeCanvasLayer->Canvases.At(LastEdge)->highLight = false;
      update();

      std::this_thread::sleep_for(std::chrono::seconds(2));
      close();
    }).detach();
  }

  private:
  bool AIPlayer1;
  bool AIPlayer2;
  Ptr<AIInterface> Player1Model;
  Ptr<AIInterface> Player2Model;
  Edge PlayerMoveEdge;
  Ptr<BoardV2> Board;
  Ptr<BoxCanvasLayer> BoxCanvasLayer;
  Ptr<EdgeCanvasLayer> EdgeCanvasLayer;
  Ptr<DotCanvasLayer> DotCanvasLayer;
  Ptr<EdgeButtonLayer> EdgeButtonLayer;
  Edge LastEdge;

  void
  setPlayerMoveEdge(Edge edge) {
    if (Board->Contains(edge)) {
      return;
    }
    if (AIPlayer1 && Board->Turn == Player1Turn) {
      return;
    }
    if (AIPlayer2 && Board->Turn == Player2Turn) {
      return;
    }
    PlayerMoveEdge = edge;
  }
};
