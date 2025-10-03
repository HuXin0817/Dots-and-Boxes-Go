#pragma once

#include <memory>
#include <string>

#include "Interface.h"
#include "L0Model.h"
#include "L1Model.h"
#include "L2Model.h"
#include "L3Model.h"
#include "L4Model.h"

enum class AIModelType { L0, L1, L2, L3, L4 };

class AIConfig {
  public:
  static std::unique_ptr<AIInterface>
  createModel(AIModelType type) {
    switch (type) {
      case AIModelType::L0:
        return std::make_unique<L0Model>();
      case AIModelType::L1:
        return std::make_unique<L1Model>();
      case AIModelType::L2:
        return std::make_unique<L2Model>();
      case AIModelType::L3:
        return std::make_unique<L3Model>();
      case AIModelType::L4:
        return std::make_unique<L4Model>();
      default:
        return std::make_unique<L4Model>();
    }
  }

  static std::string
  getModelName(AIModelType type) {
    switch (type) {
      case AIModelType::L0:
        return "L0";
      case AIModelType::L1:
        return "L1";
      case AIModelType::L2:
        return "L2";
      case AIModelType::L3:
        return "L3";
      case AIModelType::L4:
        return "L4";
      default:
        return "L4";
    }
  }

  static std::string
  getModelDescription(AIModelType type) {
    switch (type) {
      case AIModelType::L0:
        return "Basic Model - Simple Strategy";
      case AIModelType::L1:
        return "Simple Search Model - Basic Search";
      case AIModelType::L2:
        return "Intermediate Search Model - Improved Search";
      case AIModelType::L3:
        return "Advanced Search Model - Monte Carlo Search";
      case AIModelType::L4:
        return "Highest Level Model - Parallel Monte Carlo Search";
      default:
        return "Highest Level Model - Parallel Monte Carlo Search";
    }
  }

  static AIModelType
  parseModelType(const std::string& name) {
    if (name == "L0" || name == "l0") {
      return AIModelType::L0;
    } else if (name == "L1" || name == "l1") {
      return AIModelType::L1;
    } else if (name == "L2" || name == "l2") {
      return AIModelType::L2;
    } else if (name == "L3" || name == "l3") {
      return AIModelType::L3;
    } else if (name == "L4" || name == "l4") {
      return AIModelType::L4;
    } else {
      return AIModelType::L4;
    }
  }
};
