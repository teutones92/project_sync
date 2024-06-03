import 'package:project_sync/models/repository/status_code_model/status_code_model.dart';
import 'package:project_sync/models/repository/user_model/user_model.dart';
import 'package:project_sync/services/user_service/user_service.dart';
import 'package:flutter/material.dart';

import '../../core/config/bloc_config.dart';

class UserState extends UserModel {
  UserState({
    required super.id,
    required super.username,
    required super.email,
    required super.dob,
    required super.phoneNumber,
    required super.countryCode,
    required super.countryPhoneCode,
    required super.langCode,
    required super.passwordHash,
    required super.password,
    required super.userAvatarPath,
    required super.darkMode,
  });
}

class UserDataBloc extends Cubit<UserState?> {
  UserDataBloc() : super(null);

  Future<void> getUserData(BuildContext context) async {
    final messenger = ScaffoldMessenger.of(context);
    final resp = await UserService.readUserData();
    if (resp is UserModel) {
      emit(UserState(
        id: resp.id,
        username: resp.username,
        email: resp.email,
        dob: resp.dob,
        phoneNumber: resp.phoneNumber,
        countryCode: resp.countryCode,
        countryPhoneCode: resp.countryPhoneCode,
        langCode: resp.langCode,
        passwordHash: resp.passwordHash,
        password: resp.password,
        userAvatarPath: resp.userAvatarPath,
        darkMode: resp.darkMode,
      ));
      return;
    }
    if (resp is StatusCodeModel) {
      messenger.showSnackBar(
        SnackBar(
          content: Text(resp.statusCodeMessage),
        ),
      );
    }
  }

  void updateDarkModeUserData(
      {required bool value, required BuildContext context}) async {
    state!.darkMode = value;
    await UserService.updateUserData(state!).then((resp) {
      if (resp.statusCode == 200) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text("Dark Mode Updated"),
            duration: Duration(milliseconds: 500),
          ),
        );
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(resp.statusCodeMessage),
          ),
        );
      }
    });
  }
}
