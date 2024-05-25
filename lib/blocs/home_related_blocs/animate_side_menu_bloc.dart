import 'package:project_sync/core/config/bloc_config.dart';

class AnimateSideMenuBloc extends Cubit<double> {
  AnimateSideMenuBloc() : super(350);

  void animateSideMenu() => emit(state == 350 ? 0 : 350);

  void isHide(bool hide) => emit(hide ? 0 : 350);
}
