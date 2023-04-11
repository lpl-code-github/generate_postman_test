var schema = {
  "properties": {
    "date_states": {
      "items": {
        "properties": {
          "created_at": {
            "type": "string"
          },
          "id": {
            "type": [
              "string",
              "integer"
            ]
          },
          "label": {
            "type": "string"
          },
          "name": {
            "type": "string"
          },
          "short_label": {
            "type": "string"
          },
          "updated_at": {
            "type": "string"
          }
        },
        "required": [
          "id",
          "name",
          "label",
          "short_label",
          "created_at",
          "updated_at"
        ],
        "type": "object"
      },
      "type": "array"
    },
    "token_user": {
      "properties": {
        "active": {
          "type": [
            "string",
            "integer"
          ]
        },
        "created_at": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "forename": {
          "type": "string"
        },
        "full_name": {
          "type": "string"
        },
        "hidden_folders": {
          "type": "null"
        },
        "id": {
          "type": [
            "string",
            "integer"
          ]
        },
        "is_observing": {
          "type": "boolean"
        },
        "login_state": {
          "type": [
            "string",
            "integer"
          ]
        },
        "name": {
          "type": "string"
        },
        "owner_type": {
          "type": "string"
        },
        "project_order": {
          "items": {
            "properties": {
              "-1": {
                "type": "array"
              }
            },
            "required": [
              "-1"
            ],
            "type": "object"
          },
          "type": "array"
        },
        "properties": {
          "properties": {
            "dateFormat": {
              "type": "string"
            },
            "dateSymbol": {
              "type": "string"
            },
            "lang": {
              "type": "string"
            },
            "timeFormat": {
              "type": "string"
            }
          },
          "required": [
            "dateSymbol",
            "timeFormat",
            "lang",
            "dateFormat"
          ],
          "type": "object"
        },
        "surname": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "updated_at": {
          "type": "string"
        }
      },
      "required": [
        "properties",
        "owner_type",
        "hidden_folders",
        "created_at",
        "is_observing",
        "surname",
        "first_name",
        "token",
        "active",
        "updated_at",
        "project_order",
        "id",
        "forename",
        "full_name",
        "login_state",
        "name",
        "email",
        "type"
      ],
      "type": "object"
    }
  },
  "required": [
    "date_states",
    "token_user"
  ],
  "type": "object"
};

pm.test('Schema is valid', function() {
	var jsonData = pm.response.json();
	pm.expect(tv4.validate(jsonData, schema)).to.be.true;
});
