#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <memory>

#include "../common/Array.h"
#include "../model/Edge.h"
#include "BaseCanvasLayer.h"
#include "DotCanvas.h"
#include "EdgeButton.h"

class EdgeButtonLayer final : public BaseCanvasLayer {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeButtonLayer(const std::function<std::function<void()>(Edge)>& CallBackFactory,
                           QWidget* parent = nullptr);

  protected:
  void
  resizeEvent(QResizeEvent* event) override;

  private:
  Array<std::unique_ptr<EdgeButton>, Edge::Max> EdgeButtons;
};