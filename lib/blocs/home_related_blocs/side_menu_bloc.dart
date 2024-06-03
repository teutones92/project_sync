import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/models/app/side_menu_data_model/side_menu_data.dart';
import 'package:project_sync/views/0.7_settings/settings.dart';

import '../../core/config/bloc_config.dart';

class SideMenuBloc extends Cubit<int> {
  SideMenuBloc() : super(0);

  onTap(
      {required int index,
      required BuildContext context,
      required SideMenuData data}) {
    final homeBloc = context.read<HomeBloc>();
    switch (index) {
      case 6:
        NavigatorX.push(context: context, page: const Settings());
        break;
      default:
        emit(index);
        homeBloc.landingPageCtrl.jumpToPage(index);
    }
    if (MediaQuery.of(context).size.width < 1200 && index != 6) {
      context.read<AnimateSideMenuBloc>().isHide(true);
    }
  }
}
