import 'package:project_sync/blocs/users_related_blocs/user_data_bloc.dart';
import 'package:project_sync/global/http_method/http_method.dart';
import 'package:project_sync/models/repository/status_code_model/status_code_model.dart';
import 'package:project_sync/models/repository/user_model/user_model.dart';
import 'package:project_sync/services/connection_service/connection_service.dart';
import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class UserService {
  static Future readUserData() async {
    final pref = await SharedPreferences.getInstance();
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.get,
          subUrl: "user/read",
          data: {
            'id': pref.getInt('userId'),
          });
      if (resp.data['status_code'] != null) {
        return StatusCodeModel.fromJson(resp.data);
      }
      return UserModel.fromJson(resp.data);
    } on DioException catch (_) {
      return StatusCodeModel(
          statusCode: 404, statusCodeMessage: "Connection Error");
    }
  }

  static Future<StatusCodeModel> updateUserData(UserState userState) async {
    try {
      final resp = await ConnectionService.connectionHandler(
        method: HttpMethod.put,
        subUrl: "user/update",
        data: userState.toJson(),
      );
      return StatusCodeModel.fromJson(resp.data);
    } on DioException catch (_) {
      return StatusCodeModel(
          statusCode: 404, statusCodeMessage: "Connection Error");
    }
  }
}
