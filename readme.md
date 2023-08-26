# RENTX API

## Endpoints:

- **URL**: /api/v1/register
- **Method**: POST
- **Description**: Register a new car.

- **Request body**:

  ```json
  {
    "brand": "string",
    "model": "string",
    "price": "string",
    "specs": {
      "description": "string",
      "maxVelocity": "string",
      "numberOfPeople": "string",
      "fueltype": "string",
      "zeroToHundred": "string",
      "horsePower": "string",
      "engine": "string"
    }
  }
  ```

- **Response**:

  ```json
  {
    "id": "string",
    "brand": "string",
    "model": "string",
    "price": "string",
    "specs": {
      "description": "string",
      "maxVelocity": "string",
      "numberOfPeople": "string",
      "fueltype": "string",
      "zeroToHundred": "string",
      "horsePower": "string",
      "engine": "string"
    }
  }
  ```

- **URL**: /api/v1/cars
- **Method**: GET
- **Description**: List all cars.

- **Response**:

  ```json
  [
    {
      "id": "string",
      "brand": "string",
      "model": "string",
      "price": "string",
      "specs": {
        "description": "string",
        "maxVelocity": "string",
        "numberOfPeople": "string",
        "fueltype": "string",
        "zeroToHundred": "string",
        "horsePower": "string",
        "engine": "string"
      }
    }
  ]
  ```

- **URL**: /api/v1/cars/:id
- **Method**: PUT
- **Description**: Update a car.

- **Request body**:

  ```json
  {
    "brand": "string",
    "model": "string",
    "price": "string",
    "specs": {
      "description": "string",
      "maxVelocity": "string",
      "numberOfPeople": "string",
      "fueltype": "string",
      "zeroToHundred": "string",
      "horsePower": "string",
      "engine": "string"
    }
  }
  ```

- **Response**:

  ```json
  {
    "id": "string",
    "brand": "string",
    "model": "string",
    "price": "string",
    "specs": {
      "description": "string",
      "maxVelocity": "string",
      "numberOfPeople": "string",
      "fueltype": "string",
      "zeroToHundred": "string",
      "horsePower": "string",
      "engine": "string"
    }
  }
  ```

- **URL**: /api/v1/cars/:id
- **Method**: DELETE
- **Description**: Delete a car.

- **Response**:
  ```json
  { "status": "success" }
  ```
