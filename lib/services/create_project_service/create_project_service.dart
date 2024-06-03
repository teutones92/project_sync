import 'package:project_sync/global/http_method/http_method.dart';
import 'package:project_sync/services/connection_service/connection_service.dart';

class CreateProjectService {
  // Read all projects
  Future<void> readProjects() async {
    try {
      ConnectionService.connectionHandler(
          method: HttpMethod.get, subUrl: "projects/read", data: {});
    } catch (e) {
      // Handle error
    }
  }

  // Create a new project
  Future<void> createProject() async {
    try {
      ConnectionService.connectionHandler(
          method: HttpMethod.post, subUrl: "projects/create", data: {});
    } catch (e) {
      // Handle error
    }
  }

  // Update an existing project
  Future<void> updateProject() async {}

  // Delete a project
  Future<void> deleteProject() async {}
}
