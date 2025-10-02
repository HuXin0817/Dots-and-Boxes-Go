#pragma once

#include <QResizeEvent>
#include <QWidget>
#include <array>
#include <memory>

#include "../model/Edge.h"
#include "DotCanvas.h"
#include "EdgeButton.h"
#include "config.h"

class EdgeButtonLayer final : public QWidget {
  friend class MainWindow;

  Q_OBJECT

  public:
  explicit EdgeButtonLayer(const std::function<std::function<void()>(Edge)>& CallBackFactory,
                           QWidget* parent = nullptr);

  protected:
  void
  resizeEvent(QResizeEvent* event) override;

  private:
  std::array<std::unique_ptr<EdgeButton>, Edge::Max> EdgeButtons;
};