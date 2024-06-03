import 'package:flutter/material.dart';
import 'package:project_sync/core/config/bloc_config.dart';

class ProjectLeadBloc extends Cubit<UserState?> {
  ProjectLeadBloc() : super(null);

  void setProjectLead(UserState? value, BuildContext context) {
    value = context.read<UserDataBloc>().state;
    emit(value);
  }
}
