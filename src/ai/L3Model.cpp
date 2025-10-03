#include "L3Model.h"

Span<const Edge>
L3Model::BestCandidateEdges(const BoardV2& board) {
  if (auto l = SubModel.BestCandidateEdges(board); l.size() == 1) {
    return l;
  }

  ScoreMap.Reset();
  int Times = SearchTime / board.RemainStep() + 1;
  while (Times--) {
    AuxBoard.Reset(board.GetBoardV1());
    auto e = RandomChoice(SubModel.BestCandidateEdges(AuxBoard));
    AuxBoard.Add(e);
    while (AuxBoard.Gaming()) {
      AuxBoard.Add(RandomChoice(SubModel.BestCandidateEdges(AuxBoard)));
    }
    ScoreMap.Add(e, AuxBoard.Score());
  }

  return ScoreMap.Export();
}
