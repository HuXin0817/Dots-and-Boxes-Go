#pragma once

#include <QApplication>
#include <QColor>
#include <QPaintEvent>
#include <QPainter>
#include <QResizeEvent>
#include <QShowEvent>
#include <QTimer>
#include <QWidget>
#include <chrono>
#include <memory>
#include <thread>

#include "../ai/AIConfig.h"
#include "BoxCanvasLayer.h"
#include "DotCanvasLayer.h"
#include "EdgeButtonLayer.h"
#include "EdgeCanvasLayer.h"

class MainWindow : public BaseCanvasLayer {
  Q_OBJECT
  public:
  MainWindow(bool OpenAIPlayer1,
             bool OpenAIPlayer2,
             AIModelType AIPlayer1Type,
             AIModelType AIPlayer2Type,
             QWidget* parent = nullptr);

  QColor
  BackGroundColor();

  signals:
  void
  requestAIMove();

  public slots:
  void
  Add(Edge e);

  protected:
  void
  paintEvent(QPaintEvent* event) override;
  void
  resizeEvent(QResizeEvent* event) override;
  void
  showEvent(QShowEvent* event) override;

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
};