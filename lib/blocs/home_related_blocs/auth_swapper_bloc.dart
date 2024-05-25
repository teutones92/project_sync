import 'package:project_sync/core/config/bloc_config.dart';
import 'package:flutter/material.dart';

class AuthSwapperBloc extends Cubit<bool> {
  AuthSwapperBloc() : super(true);

  void swap(BuildContext context) {
    final loginBloc = context.read<LoginBloc>();
    final registerBloc = context.read<RegisterBloc>();
    loginBloc.emailController.clear();
    loginBloc.passwordController.clear();
    registerBloc.firstNameController.clear();
    registerBloc.lastNameController.clear();
    registerBloc.emailController.clear();
    registerBloc.passwordController.clear();
    emit(!state);
  }
}
