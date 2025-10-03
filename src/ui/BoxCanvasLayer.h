#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Box.h"
#include "BaseCanvasLayer.h"
#include "BoxCanvas.h"
#include "DotCanvas.h"
#include "EdgeCanvas.h"

class BoxCanvasLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit BoxCanvasLayer(QWidget* parent = nullptr);

  protected:
  void
  resizeEvent(QResizeEvent* event) override;

  private:
  Array<std::unique_ptr<BoxCanvas>, Box::Max> BoxCanvases;
};