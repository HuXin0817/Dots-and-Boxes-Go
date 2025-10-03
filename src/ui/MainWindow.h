#pragma once

#include <QApplication>
#include <QColor>
#include <QTimer>
#include <chrono>
#include <memory>
#include <thread>

#include "../ai/AIConfig.h"
#include "BoxCanvasLayer.h"
#include "DotCanvasLayer.h"
#include "EdgeButtonLayer.h"
#include "EdgeCanvasLayer.h"

class MainWindow final : public BaseCanvasLayer {
  Q_OBJECT
  public:
  MainWindow(bool OpenAIPlayer1,
             bool OpenAIPlayer2,
             AIModelType AIPlayer1Type,
             AIModelType AIPlayer2Type,
             QWidget* parent = nullptr)
      : BaseCanvasLayer(parent),
        OpenAIPlayer1(OpenAIPlayer1),
        OpenAIPlayer2(OpenAIPlayer2),
        AIPlayer1(AIConfig::createModel(AIPlayer1Type)),
        AIPlayer2(AIConfig::createModel(AIPlayer2Type)) {
    resize(WindowSize, WindowSize);
    setMinimumSize(WindowSize, WindowSize);

    board = std::make_unique<BoardV2>();
    boxLayer = std::make_unique<BoxCanvasLayer>(this);
    edgeLayer = std::make_unique<EdgeCanvasLayer>(this);
    dotLayer = std::make_unique<DotCanvasLayer>(this);
    std::function<std::function<void()>(Edge)> CallBackFactory = [this](Edge e) {
      return [e, this] { return setPlayerMoveEdge(e); };
    };
    edgeButtonLayer = std::make_unique<EdgeButtonLayer>(CallBackFactory, this);
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
  Add(Edge e) {
    auto Turn = board->Turn;
    if (board->NowStep() > 0) {
      edgeLayer->EdgeCanvases.At(lastEdge)->highLight = false;
    }
    edgeLayer->EdgeCanvases.At(e)->state = StateFromTurn(Turn);
    edgeLayer->EdgeCanvases.At(e)->raise();

    for (auto box : EdgeBoxMapper::EdgeNearBoxes.At(e)) {
      int c = 0;
      for (auto nearEdge : EdgeBoxMapper::BoxNearEdges.At(box)) {
        if (board->Contains(nearEdge)) {
          c++;
        }
      }
      if (c == 3) {
        boxLayer->BoxCanvases.At(box)->state = StateFromTurn(Turn);
      }
    }

    board->Add(e);
    lastEdge = e;
    update();
    QApplication::beep();
  }

  protected:
  void
  paintEvent(QPaintEvent* event) override {
    QPainter painter(this);
    painter.fillRect(rect(), Color());
  }

  void
  resizeEvent(QResizeEvent* event) override {
    QWidget::resizeEvent(event);

    int x = (width() - WindowSize) / 2;
    int y = (height() - WindowSize) / 2;

    boxLayer->move(x, y);
    edgeLayer->move(x, y);
    dotLayer->move(x, y);
  }

  void
  showEvent(QShowEvent* event) override {
    QWidget::showEvent(event);

    std::thread([this] {
      while (board->Gaming()) {
        if (OpenAIPlayer1 && board->Turn == Player1Turn) {
          PlayerMoveEdge = RandomChoice(AIPlayer1->BestCandidateEdges(*board));
        } else if (OpenAIPlayer2 && board->Turn == Player2Turn) {
          PlayerMoveEdge = RandomChoice(AIPlayer2->BestCandidateEdges(*board));
        } else {
          PlayerMoveEdge = -1;
          while (PlayerMoveEdge == -1) {
            std::this_thread::yield();
          }
        }
        Add(PlayerMoveEdge);

        int playerId = board->Turn == Player1Turn ? 1 : 2;
        int step = board->NowStep();

        printf("| Step %d | Player %d Move (%d, %d) -> (%d, %d) | Score %d : %d |\n",
               step,
               playerId,
               PlayerMoveEdge.Dot1().X(),
               PlayerMoveEdge.Dot1().Y(),
               PlayerMoveEdge.Dot2().X(),
               PlayerMoveEdge.Dot2().Y(),
               board->Player1Score,
               board->Player2Score);
      }

      if (board->Player1Score > board->Player2Score) {
        printf("| Player 1 Win! |\n");
      } else if (board->Player2Score > board->Player1Score) {
        printf("| Player 2 Win! |\n");
      } else {
        printf("| Draw! |\n");
      }

      std::this_thread::sleep_for(std::chrono::seconds(2));
      edgeLayer->EdgeCanvases.At(lastEdge)->highLight = false;
      update();

      std::this_thread::sleep_for(std::chrono::seconds(2));
      close();
    }).detach();
  }

  private:
  bool OpenAIPlayer1;
  bool OpenAIPlayer2;
  std::unique_ptr<AIInterface> AIPlayer1;
  std::unique_ptr<AIInterface> AIPlayer2;
  Edge PlayerMoveEdge;
  std::unique_ptr<BoardV2> board;
  std::unique_ptr<BoxCanvasLayer> boxLayer;
  std::unique_ptr<EdgeCanvasLayer> edgeLayer;
  std::unique_ptr<DotCanvasLayer> dotLayer;
  std::unique_ptr<EdgeButtonLayer> edgeButtonLayer;
  Edge lastEdge;

  void
  setPlayerMoveEdge(Edge e) {
    if (board->Contains(e)) {
      return;
    }
    if (OpenAIPlayer1 && board->Turn == Player1Turn) {
      return;
    }
    if (OpenAIPlayer2 && board->Turn == Player2Turn) {
      return;
    }
    PlayerMoveEdge = e;
  }
};
