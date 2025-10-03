#include "L4Model.h"

Span<const Edge>
L4Model::BestCandidateEdges(const BoardV2& board) {
  if (auto l = L2Model().BestCandidateEdges(board); l.size() == 1) {
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
