#pragma once

#include "Interface.h"
#include "L3Model.h"

class L4Model final : public AIInterface {
  public:
  static constexpr int SubModelSearchTime = 1000;
  static constexpr int SearchGroupNumber = 100;

  L4Model() = default;

  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto edges = L2Model().BestCandidateEdges(board); edges.Size() == 1) {
      return edges;
    }

    thread_local L3Model model(SubModelSearchTime);
    EdgeScoreMap result;

#pragma omp parallel for
    for (int i = 0; i < SearchGroupNumber; ++i) {
      model.BestCandidateEdges(board);
#pragma omp critical
      {
        result.Add(model.ScoreMap);
      }
    }

    return result.Export();
  }
};
