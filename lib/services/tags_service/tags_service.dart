import 'package:project_sync/global/http_method/http_method.dart';
import 'package:project_sync/services/connection_service/connection_service.dart';

class TagsService {
  // Future<List<Tag>> getTags() async {
  //   // Fetch tags from the server
  //   return tags;
  // }

  static Future<void> readTags() async {
    // Read tags from the local storage
    final resp = await ConnectionService.connectionHandler(
        method: HttpMethod.get, subUrl: "project_tags/read", data: {});
    print(resp);
  }
}
