#pragma once

#include "../common/Vector.h"
#include "Interface.h"
#include "L3Model.h"

class L4Model final : public AIInterface {
  public:
  static constexpr int SubModelSearchTime = 1000;

  explicit L4Model(int GroupNumber = 100) : SearchResults(GroupNumber) {
  }

  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto edges = L2Model().BestCandidateEdges(board); edges.Size() == 1) {
      return edges;
    }

    thread_local L3Model model(SubModelSearchTime);

#pragma omp parallel for
    for (auto& searchResult : SearchResults) {
      model.BestCandidateEdges(board);
      searchResult = model.ScoreMap;
    }

    EdgeScoreMap result;
    for (const auto& searchResult : SearchResults) {
      result.Add(searchResult);
    }

    return result.Export();
  }

  private:
  Vector<EdgeScoreMap> SearchResults;
};
