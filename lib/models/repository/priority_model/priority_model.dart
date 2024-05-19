class PriorityModel {
  final int id;
  final String name;
  final String description;
  final String color;

  PriorityModel(
      {required this.id,
      required this.name,
      required this.description,
      required this.color});

  factory PriorityModel.fromJson(Map<String, dynamic> json) {
    return PriorityModel(
      id: json['id'],
      name: json['priority_name'],
      description: json['priority_description'],
      color: json['priority_rgb_color'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'priority_name': name,
      'priority_description': description,
      'priority_rgb_color': color,
    };
  }
}
