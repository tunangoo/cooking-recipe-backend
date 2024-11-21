# cooking-recipe-backend
URL: https://cooking-recipe-backend.onrender.com

## Project Overview
We are looking to build a simple cooking recipe web application where users can search and view recipes. The application should allow users to browse recipes based on ingredients and cuisine. The recipes should be displayed with a picture, ingredients list, and step-by-step instructions.

## Technologies
- **Backend**: Golang (Gin Framework)
- **Database**: MySQL
- **Deployment**: Render.com, Railway.app

## Key Features
- Recipe search and filter options
- Recipe detail page with ingredients list, instructions, and photo

## API Endpoints
## 1. **GET /api/recipe/all**
- **Description**: Retrieves a list of all recipes.
- **Usage**: This endpoint returns a complete list of recipes stored in the database.
- **Example Request**:
    - GET https://cooking-recipe-backend.onrender.com/api/recipe/all
- **Response**: Returns an array of all recipes.

## 2. **GET /api/recipe/{id}**
- **Description**: Retrieves a specific recipe by its ID.
- **Usage**: This endpoint fetches detailed information for a single recipe identified by the `id` parameter.
- **Example Request**:
    - GET https://cooking-recipe-backend.onrender.com/api/recipe/1
- **Response**: Returns details of the recipe with the specified `id`, such as title, ingredients, instructions, and photo.

## 3. **GET /api/recipe/search**
- **Description**: Searches for recipes based on provided query parameters (e.g., name or ingredient).
- **Usage**: This endpoint allows users to search recipes by parameters like name or ingredient.
- **Example Request**:
    - GET https://cooking-recipe-backend.onrender.com/api/recipe/search?q=chicken
- **Response**: Returns a list of recipes that match the search criteria.