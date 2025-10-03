#pragma once

#include "../common/List.h"
#include "../common/Span.h"
#include "L1Model.h"

class L2Model final : public AIInterface {
  public:
  L2Model() = default;

  Span<Edge>
  BestCandidateEdges(const BoardV2& board) override {
    if (auto edges = SubModel.BestCandidateEdges(board);
        !SubModel.SubModel.EnemyUnscoreableEdges.Empty()) {
      return edges;
    }

    SearchEdges.Clear();
    int maxScore = -(Box::Max + 1);
    for (Edge emptyEdge : board.EmptyEdges()) {
      AuxBoard.Reset(board.GetBoardV1());
      AuxBoard.Add(emptyEdge);
      while (AuxBoard.Gaming()) {
        Edge edge = SubModel.BestCandidateEdges(AuxBoard).At(0);
        assert(board.MaxCount(edge) > 1);
        AuxBoard.Add(edge);
      }
      if (int score = AuxBoard.Score(); score > maxScore) {
        maxScore = score;
        SearchEdges.Reset(emptyEdge);
      } else if (score == maxScore) {
        SearchEdges.Append(emptyEdge);
      }
    }

    return SearchEdges.Export();
  }

  private:
  L1Model SubModel;
  BoardV2 AuxBoard;
  List<Edge, Edge::Max> SearchEdges;
};
