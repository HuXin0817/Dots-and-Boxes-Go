#include "MainWindow.h"

MainWindow::MainWindow(bool OpenAIPlayer1,
                       bool OpenAIPlayer2,
                       AIModelType AIPlayer1Type,
                       AIModelType AIPlayer2Type,
                       QWidget* parent)
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
  edgeButtonLayer = std::make_unique<EdgeButtonLayer>(
      [=, this](Edge e) {
        return [=, this] {
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
        };
      },
      this);
}

void
MainWindow::Add(Edge e) {
  auto Turn = board->Turn;
  if (board->NowStep() > 0) {
    edgeLayer->EdgeCanvases[lastEdge]->highLight = false;
  }
  edgeLayer->EdgeCanvases[e]->state = StateFromTurn(Turn);
  edgeLayer->EdgeCanvases[e]->raise();

  for (auto box : EdgeBoxMapper::EdgeNearBoxes[e]) {
    int c = 0;
    for (auto nearEdge : EdgeBoxMapper::BoxNearEdges[box]) {
      if (board->Contains(nearEdge)) {
        c++;
      }
    }
    if (c == 3) {
      boxLayer->BoxCanvases[box]->state = StateFromTurn(Turn);
    }
  }

  board->Add(e);
  lastEdge = e;
  update();
  QApplication::beep();
}

QColor
MainWindow::BackGroundColor() {
  if (isDarkMode()) {
    return {43, 43, 43, 255};
  } else {
    return {242, 242, 242, 255};
  }
}

void
MainWindow::paintEvent(QPaintEvent* event) {
  QPainter painter(this);
  painter.fillRect(rect(), BackGroundColor());
}

void
MainWindow::resizeEvent(QResizeEvent* event) {
  QWidget::resizeEvent(event);

  int x = (width() - WindowSize) / 2;
  int y = (height() - WindowSize) / 2;

  boxLayer->move(x, y);
  edgeLayer->move(x, y);
  dotLayer->move(x, y);
}

void
MainWindow::showEvent(QShowEvent* event) {
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
             PlayerMoveEdge.dot1().X(),
             PlayerMoveEdge.dot1().Y(),
             PlayerMoveEdge.dot2().X(),
             PlayerMoveEdge.dot2().Y(),
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
    edgeLayer->EdgeCanvases[lastEdge]->highLight = false;
    update();

    std::this_thread::sleep_for(std::chrono::seconds(2));
    close();
  }).detach();
}
