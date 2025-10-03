#pragma once

#include <omp.h>

#include <thread>

#include "Interface.h"
#include "L3Model.h"

class L4Model final : public AIInterface {
  public:
  static constexpr int SubModelSearchTime = 1000;

  explicit L4Model(int GroupNumber = 100) : GroupNumber(GroupNumber) {
  }

  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto l = L2Model().BestCandidateEdges(board); l.Size() == 1) {
      return l;
    }

    std::vector<EdgeScoreMap> SearchResults(GroupNumber);
    thread_local L3Model model(SubModelSearchTime);

#pragma omp parallel for
    for (auto& r : SearchResults) {
      model.BestCandidateEdges(board);
      r = model.ScoreMap;
    }

    EdgeScoreMap result;
    for (const auto& r : SearchResults) {
      result += r;
    }

    return result.Export();
  }

  private:
  int GroupNumber;
};