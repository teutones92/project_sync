import 'package:project_sync/core/config/bloc_config.dart';

import 'package:project_sync/core/extensions.dart';
import 'package:project_sync/services/auth_service/auth_service.dart';
import 'package:project_sync/views/0.0_authentication/authentication.dart';
import 'package:flutter/material.dart';

abstract class HomeEvent {}

class HomeBloc extends Bloc<HomeEvent, bool> {
  HomeBloc() : super(false);

  PageController landingPageCtrl = PageController();

  Future<void> logOut(BuildContext context) async {
    await AuthService.logout().then((value) {
      if (value.statusCode == 200) {
        NavigatorX.pushAndRemoveUntil(
          context: context,
          page: const Authentication(),
        ).then((_) {
          context.read<DarkModeBloc>().setSystemMode();
        });
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(value.statusCodeMessage),
          ),
        );
      }
    });
  }
}
