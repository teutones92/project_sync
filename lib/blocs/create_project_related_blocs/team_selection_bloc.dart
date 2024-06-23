import 'package:project_sync/core/config/bloc_config.dart';

class TeamSelectionBloc extends Cubit<int?> {
  TeamSelectionBloc() : super(null);

  void setTeamSelection(int value) {
    if (state == value) {
      resetTeamSelection();
    } else {
      emit(value);
    }
  }

  void resetTeamSelection() => emit(null);
}
