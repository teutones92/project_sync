import 'package:flutter/material.dart';
import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/views/0.7_settings/settings.dart';

import '../core/config/bloc_config.dart';

class SideMenuBloc extends Bloc {
  SideMenuBloc() : super(null);

  onTap({required int index, required BuildContext context}) {
    switch (index) {
      case 0:
        break;
      case 1:
        break;
      case 2:
        break;
      case 5:
        NavigatorX.push(context: context, page: const Settings());
        break;
    }
  }
}
