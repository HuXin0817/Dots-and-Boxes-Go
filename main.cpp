#include <QApplication>
#include <QString>
#include <QStyleFactory>

#include "src/ai/AIConfig.h"
#include "src/ui/MainWindow.h"

int
main(int argc, char* argv[]) {
  QApplication app(argc, argv);

  QApplication::setApplicationName("Dots and Boxes");
  QApplication::setApplicationVersion("1.0");
  QApplication::setOrganizationName("Dots and Boxes");
  QApplication::setStyle(QStyleFactory::create("Fusion"));

  AIModelType player1Model = AIModelType::L4;
  AIModelType player2Model = AIModelType::L4;
  bool aiPlayer1 = true;
  bool aiPlayer2 = true;

  for (int i = 1; i < argc; i++) {
    std::string arg = argv[i];
    if (arg == "--player1" && i + 1 < argc) {
      player1Model = AIConfig::parseModelType(argv[++i]);
    } else if (arg == "--player2" && i + 1 < argc) {
      player2Model = AIConfig::parseModelType(argv[++i]);
    } else if (arg == "--ai-player1") {
      aiPlayer1 = true;
    } else if (arg == "--human-player1") {
      aiPlayer1 = false;
    } else if (arg == "--ai-player2") {
      aiPlayer2 = true;
    } else if (arg == "--human-player2") {
      aiPlayer2 = false;
    } else if (arg == "--help" || arg == "-h") {
      printf("Dots and Boxes - AI Model Configuration\n");
      printf("Usage: %s [options]\n", argv[0]);
      printf("Options:\n");
      printf("  --player1 <model>  Set AI model for player 1 (L0, L1, L2, L3, L4)\n");
      printf("  --player2 <model>  Set AI model for player 2 (L0, L1, L2, L3, L4)\n");
      printf("  --ai-player1       Set player 1 as AI (default)\n");
      printf("  --human-player1    Set player 1 as human\n");
      printf("  --ai-player2       Set player 2 as AI (default)\n");
      printf("  --human-player2    Set player 2 as human\n");
      printf("  --help, -h         Show this help message\n");
      printf("\nAvailable AI Models:\n");
      printf("  L0: %s\n", AIConfig::getModelDescription(AIModelType::L0).c_str());
      printf("  L1: %s\n", AIConfig::getModelDescription(AIModelType::L1).c_str());
      printf("  L2: %s\n", AIConfig::getModelDescription(AIModelType::L2).c_str());
      printf("  L3: %s\n", AIConfig::getModelDescription(AIModelType::L3).c_str());
      printf("  L4: %s\n", AIConfig::getModelDescription(AIModelType::L4).c_str());
      return 0;
    }
  }

  printf("Starting game with player configuration:\n");
  printf("  Player 1: %s", aiPlayer1 ? "AI" : "Human");
  if (aiPlayer1) {
    printf(" (%s - %s)",
           AIConfig::getModelName(player1Model).c_str(),
           AIConfig::getModelDescription(player1Model).c_str());
  }
  printf("\n");
  printf("  Player 2: %s", aiPlayer2 ? "AI" : "Human");
  if (aiPlayer2) {
    printf(" (%s - %s)",
           AIConfig::getModelName(player2Model).c_str(),
           AIConfig::getModelDescription(player2Model).c_str());
  }
  printf("\n\n");

  auto mainWindow = std::make_unique<MainWindow>(aiPlayer1, aiPlayer2, player1Model, player2Model);
  mainWindow->show();

  return QApplication::exec();
}