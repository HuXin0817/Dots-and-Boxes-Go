#pragma once

#include <QApplication>
#include <QColor>
#include <QWidget>

#include "../model/ScoreMap.h"

class BaseCanvas : public QWidget {
  public:
  explicit BaseCanvas(QWidget* parent = nullptr) : QWidget(parent) {
  }

  virtual QColor
  Color() const {
    return {};
  }

  enum class State {
    Free,
    Player1Occupy,
    Player2Occupy,
  };

  static State
  StateFromTurn(bool Turn) {
    return Turn == Player1Turn ? State::Player1Occupy : State::Player2Occupy;
  }

  static bool
  isDarkMode() {
    QPalette palette = QApplication::palette();
    QColor windowColor = palette.color(QPalette::Window);

    return windowColor.lightness() < 128;
  }

  static constexpr int R = 8;
};
