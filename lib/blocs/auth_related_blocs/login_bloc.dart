import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/core/extensions.dart';

import 'package:project_sync/services/auth_service/auth_service.dart';
import 'package:project_sync/views/0.1_home/home.dart';
import 'package:flutter/material.dart';

abstract class LoginEvent {}

class LoginBloc extends Bloc<LoginEvent, bool> {
  LoginBloc() : super(false);

  ValueNotifier<bool> isLoading = ValueNotifier(false);

  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  GlobalKey<FormState> emailFormKey = GlobalKey<FormState>();
  GlobalKey<FormState> passwordFormKey = GlobalKey<FormState>();

  Future<void> login(BuildContext context, {bool fromLogin = false}) async {
    if (fromLogin) {
      final emailVal = emailFormKey.currentState!.validate();
      final passVal = passwordFormKey.currentState!.validate();
      if (!emailVal || !passVal) {
        return;
      }
    }
    isLoading.value = true;
    await AuthService.login({
      'email': emailController.text,
      'password': passwordController.text,
    }).then((value) {
      if (value.statusCode == 200) {
        _clearAll();
        NavigatorX.pushAndRemoveUntil(
          context: context,
          page: const Home(),
        );
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(value.statusCodeMessage),
          ),
        );
      }
    });
    isLoading.value = false;
  }

  String? emailValidator(String? p1) {
    if (p1 == null || p1.isEmpty) {
      return 'Email is required';
    }
    return null;
  }

  String? passwordValidator(String? p1) {
    if (p1 == null || p1.isEmpty) {
      return 'Password is required';
    }
    return null;
  }

  void _clearAll() {
    emailController.clear();
    passwordController.clear();
  }
}
