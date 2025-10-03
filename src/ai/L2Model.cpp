#include "L2Model.h"

Span<const Edge>
L2Model::BestCandidateEdges(const BoardV2& board) {
  if (auto l = SubModel.BestCandidateEdges(board);
      !SubModel.SubModel.EnemyUnscoreableEdges.Empty()) {
    return l;
  }

  SearchEdges.Clear();
  int maxs = -(Box::Max + 1);
  for (auto e : board.EmptyEdges()) {
    AuxBoard.Reset(board.GetBoardV1());
    AuxBoard.Add(e);
    while (AuxBoard.Gaming()) {
      auto edge = SubModel.BestCandidateEdges(AuxBoard).At(0);
      assert(board.MaxCount(edge) > 1);
      AuxBoard.Add(edge);
    }
    if (int s = AuxBoard.Score(); s > maxs) {
      maxs = s;
      SearchEdges.Reset(e);
    } else if (s == maxs) {
      SearchEdges.Append(e);
    }
  }

  return SearchEdges.Export();
}
