import 'package:flutter/material.dart';
import 'package:project_sync/models/repository/priority_model/priority_model.dart';
import 'package:project_sync/models/repository/status_code_model/status_code_model.dart';
import 'package:project_sync/services/priority_service/priority_service.dart';
import '../../core/config/bloc_config.dart';

class PrioritySelectedBloc extends Cubit<int?> {
  PrioritySelectedBloc() : super(null);

  void selectPriority(int index, PriorityModel item) {
    if (state == index) {
      emit(null);
    } else {
      emit(index);
    }
  }
}

class PriorityBloc extends Cubit<List<PriorityModel>> {
  PriorityBloc() : super([]);

  Future<List<PriorityModel>> getPriorities(BuildContext context) async {
    final List<PriorityModel> priorities = [];
    await PriorityService.getPriorities().then((val) {
      if (val is StatusCodeModel) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(val.statusCodeMessage),
            duration: const Duration(milliseconds: 500),
          ),
        );
      } else {
        for (var item in val) {
          priorities.add(PriorityModel.fromJson(item));
        }
        emit(priorities);
      }
    });
    return priorities;
  }
}
