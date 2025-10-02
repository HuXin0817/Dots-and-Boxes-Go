#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <array>
#include <memory>

#include "../model/Box.h"
#include "BoxCanvas.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"
#include "config.h"

class BoxLayer final : public QWidget {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit BoxLayer(QWidget* parent = nullptr);

  protected:
  void
  resizeEvent(QResizeEvent* event) override;

  private:
  std::array<std::unique_ptr<BoxCanvas>, Box::Max> BoxCanvases;
};