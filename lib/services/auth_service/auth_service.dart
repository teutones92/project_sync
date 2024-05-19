import 'package:project_sync/global/http_method/http_method.dart';
import 'package:project_sync/models/repository/status_code_model/status_code_model.dart';
import 'package:project_sync/services/connection_service/connection_service.dart';
import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AuthService {
  static const _authUrlSubPath = 'auth/';

  static Future<StatusCodeModel> signUp(Map<String, dynamic> data) async {
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.post,
          subUrl: "${_authUrlSubPath}signup",
          data: data);
      return StatusCodeModel.fromJson(resp.data);
    } on DioException catch (_) {
      return StatusCodeModel(
          statusCode: 404, statusCodeMessage: "Connection Error");
    }
  }

  static Future<StatusCodeModel> login(Map<String, String> data) async {
    final pref = await SharedPreferences.getInstance();
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.post,
          subUrl: "${_authUrlSubPath}login",
          data: data);
      if (resp.data["status_code"] != null) {
        return StatusCodeModel.fromJson(resp.data);
      } else {
        await pref.setString('token', resp.data['token']);
        await pref.setInt('userId', resp.data['user_id']);
        await pref.setBool('isLoggedIn', true);
        return StatusCodeModel(
            statusCode: 200, statusCodeMessage: "Login Successful");
      }
    } on DioException catch (_) {
      return StatusCodeModel(
          statusCode: 404, statusCodeMessage: "Connection Error");
    }
  }

  static Future<StatusCodeModel> logout() async {
    final pref = await SharedPreferences.getInstance();
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.post,
          subUrl: "${_authUrlSubPath}logout",
          data: {});
      final statusCode = StatusCodeModel.fromJson(resp.data);
      if (statusCode.statusCode == 200) {
        pref.clear();
      }
      return statusCode;
    } on DioException catch (_) {
      return StatusCodeModel(
          statusCode: 404, statusCodeMessage: "Connection Error");
    }
  }

  static Future<StatusCodeModel> deleteAccount() async {
    final pref = await SharedPreferences.getInstance();
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.delete,
          subUrl: "${_authUrlSubPath}delete_account",
          data: {});
      final status = StatusCodeModel.fromJson(resp.data);
      if (status.statusCode == 200) {
        pref.clear();
      }
      return status;
    } on DioException catch (_) {
      return StatusCodeModel(
          statusCode: 404, statusCodeMessage: "Connection Error");
    }
  }
}
