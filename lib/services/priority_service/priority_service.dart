import 'package:project_sync/global/http_method/http_method.dart';
import 'package:project_sync/models/repository/status_code_model/status_code_model.dart';

import '../connection_service/connection_service.dart';

class PriorityService {
  static Future getPriorities() async {
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.get, subUrl: 'priority/read', data: {});
      if (resp.data is Map) {
        return StatusCodeModel.fromJson(resp.data);
      } else {
        return resp.data;
      }
    } catch (_) {
      return StatusCodeModel(
          statusCode: 500, statusCodeMessage: 'Server Error');
    }
  }
}
