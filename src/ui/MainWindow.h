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

#include "../ai/L4Model.h"
#include "BoxLayer.h"
#include "ButtonLayer.h"
#include "DotLayer.h"
#include "EdgeLayer.h"

class MainWindow : public QWidget {
  Q_OBJECT
  public:
  MainWindow(bool OpenAIPlayer1,
             bool OpenAIPlayer2,
             AIInterface* AIPlayer1,
             AIInterface* AIPlayer2,
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
  AIInterface* AIPlayer1;
  AIInterface* AIPlayer2;
  Edge PlayerMoveEdge;
  std::unique_ptr<BoardV2> board;
  std::unique_ptr<BoxLayer> boxLayer;
  std::unique_ptr<EdgeLayer> edgeLayer;
  std::unique_ptr<DotLayer> dotLayer;
  std::unique_ptr<EdgeButtonLayer> edgeButtonLayer;
  Edge lastEdge;
};