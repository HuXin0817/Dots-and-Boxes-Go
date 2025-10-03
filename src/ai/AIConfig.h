#pragma once

#include <memory>
#include <string>

#include "Interface.h"

enum class AIModelType { L0, L1, L2, L3, L4 };

class AIConfig {
  public:
  static std::unique_ptr<AIInterface>
  createModel(AIModelType type);
  static std::string
  getModelName(AIModelType type);
  static std::string
  getModelDescription(AIModelType type);
  static AIModelType
  parseModelType(const std::string& name);
};
