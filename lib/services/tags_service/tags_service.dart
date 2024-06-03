import 'package:dio/dio.dart';
import 'package:project_sync/global/http_method/http_method.dart';
import 'package:project_sync/models/repository/status_code_model/status_code_model.dart';
import 'package:project_sync/models/repository/tags_model/tags_model.dart';
import 'package:project_sync/services/connection_service/connection_service.dart';

class TagsService {
  static Future<dynamic> readTags() async {
    // Read tags from the local storage
    try {
      final resp = await ConnectionService.connectionHandler(
          method: HttpMethod.get, subUrl: "project_tags/read", data: {});
      if (resp.data is Map) {
        return StatusCodeModel.fromJson(resp.data);
      } else {
        final List<TagsModel> tags = [];
        for (var tag in resp.data) {
          tags.add(TagsModel.fromJson(tag));
        }
        return tags;
      }
    } on DioException catch (e) {
      return StatusCodeModel.fromJson(e.response!.data);
    }
  }
}
