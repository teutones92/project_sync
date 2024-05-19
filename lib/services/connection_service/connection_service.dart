import 'package:project_sync/global/http_method/http_method.dart';
import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ConnectionService {
  static const String _baseUrl = 'http://localhost:8080';
  static const String _apiVersion = '/v1/';
  static const String _apiUrl = '$_baseUrl$_apiVersion';

  static Future<Response> connectionHandler(
      {required String method,
      required String subUrl,
      required Map data}) async {
    final pref = await SharedPreferences.getInstance();
    final Map<String, String> header = {
      'Content-Type': 'application/json',
      'Authorization': pref.getString('token') ?? '',
    };
    final dio = Dio(
      BaseOptions(
        headers: header,
      ),
    );
    switch (method) {
      case HttpMethod.get:
        return await dio.get(_apiUrl + subUrl, data: data);
      case HttpMethod.post:
        return await dio.post(_apiUrl + subUrl, data: data);
      case HttpMethod.put:
        return await dio.put(_apiUrl + subUrl, data: data);
      case HttpMethod.patch:
        return await dio.patch(_apiUrl + subUrl, data: data);
      case HttpMethod.delete:
        return await dio.delete(_apiUrl + subUrl, data: data);
      default:
        throw Exception('Method not found or URL not found');
    }
  }
}
