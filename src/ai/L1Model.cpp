#include "L1Model.h"

std::span<const Edge>
L1Model::BestCandidateEdges(const BoardV2& board) {
  if (auto l = SubModel.BestCandidateEdges(board);
      !SubModel.EnemyUnscoreableEdges.Empty() || !SubModel.ScoreableEdges.Empty()) {
    return l;
  }

  int mins = Box::Max + 1;
  auto& Candidate = SubModel.EnemyUnscoreableEdges;
  assert(Candidate.Empty());

  for (auto e : board.EmptyEdges()) {
    AuxBoard.Reset(board.GetBoardV1());
    int v = AuxBoard.Add(e);
    assert(v == 0);

    int s = AuxBoard.MaxObtainableScore(mins);
    if (s < mins) {
      mins = s;
      Candidate.Reset(e);
    } else if (s == mins) {
      Candidate.Append(e);
    }
  }

  return Candidate.Export();
}
