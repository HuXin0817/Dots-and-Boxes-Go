#pragma once

#include "../model/Box.h"
#include "BoxCanvas.h"

static constexpr int BoardWindowSize = Box::Size * EdgeCanvas::B;
static constexpr int WindowSize = (Box::Size - 1) * EdgeCanvas::B + 2 * BoxCanvas::A + 80;
