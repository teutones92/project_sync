import 'package:project_sync/core/config/bloc_config.dart';
import 'package:project_sync/global/env.dart';
import 'package:project_sync/services/auth_service/auth_service.dart';
import 'package:flutter/material.dart';

abstract class RegisterBlocEvent {}

class RegisterBloc extends Bloc<RegisterBlocEvent, bool> {
  RegisterBloc() : super(false);

  TextEditingController firstNameController = TextEditingController();
  TextEditingController lastNameController = TextEditingController();
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  GlobalKey<FormState> userNameFromKey = GlobalKey<FormState>();
  GlobalKey<FormState> emailFromKey = GlobalKey<FormState>();
  GlobalKey<FormState> passwordFromKey = GlobalKey<FormState>();

  ValueNotifier<bool> isLoading = ValueNotifier<bool>(false);

  Future<void> register(BuildContext context) async {
    final emailVal = emailFromKey.currentState!.validate();
    final passVal = passwordFromKey.currentState!.validate();
    final userNameVal = userNameFromKey.currentState!.validate();
    if (!emailVal || !passVal || !userNameVal) {
      return;
    }
    isLoading.value = true;
    await AuthService.signUp({
      'username': '${firstNameController.text} ${lastNameController.text}',
      'email': emailController.text,
      'password': passwordController.text,
      'dark_mode': Theme.of(context).brightness.name == Brightness.dark.name
          ? true
          : false,
    }).then((value) async {
      ScaffoldMessenger.of(context)
          .showSnackBar(
            SnackBar(
              content: Text(value.statusCodeMessage),
              duration: const Duration(milliseconds: 500),
            ),
          )
          .closed
          .then((_) async {
        if (value.statusCode == 200) {
          final loginBLoc = context.read<LoginBloc>();
          loginBLoc.emailController.text = emailController.text;
          loginBLoc.passwordController.text = passwordController.text;
          _clearAll();
          await loginBLoc.login(context);
        }
        isLoading.value = false;
      });
    });
    isLoading.value = false;
  }

  String? userNameValidator(String? p1) {
    if (p1 == null || p1.isEmpty) {
      return 'Username is required';
    }
    return null;
  }

  String? emailValidator(String? p1) {
    if (p1 == null || p1.isEmpty) {
      return 'Email is required';
    }
    if (!emailRegExp.hasMatch(p1)) {
      return 'Invalid email';
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
    firstNameController.clear();
    lastNameController.clear();
    emailController.clear();
    passwordController.clear();
  }
}
