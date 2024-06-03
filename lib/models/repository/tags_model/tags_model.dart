class TagsModel {
  final int? id;
  final int? projectId;
  final String tagName;

  TagsModel({required this.id, this.projectId, required this.tagName});

  factory TagsModel.fromJson(Map<String, dynamic> json) {
    return TagsModel(
      id: json['id'],
      projectId: json['project_id'],
      tagName: json['tag_name'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'project_id': projectId,
      'tag_name': tagName,
    };
  }
}
