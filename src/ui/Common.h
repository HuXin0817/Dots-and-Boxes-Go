#pragma once

#include <QApplication>
#include <QColor>
#include <QPalette>

#include "../model/ScoreMap.h"

enum class State {
  Free,
  Player1Occupy,
  Player2Occupy,
};

inline State
StateFromTurn(bool Turn) {
  if (Turn == Player1Turn) {
    return State::Player1Occupy;
  } else {
    return State::Player2Occupy;
  }
}

inline bool
isDarkMode() {
  QPalette palette = QApplication::palette();
  QColor windowColor = palette.color(QPalette::Window);

  return windowColor.lightness() < 128;
}