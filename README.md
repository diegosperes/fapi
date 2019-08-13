# FAPI - The faster API.

To insert a new model:

-> (POST) http://{host}/{api-name}/{model}

To filter a model by attribute:

-> (GET) http://{host}/{api-name}/{model}?{attribute-name}={attribute-value}

To get any model by id:

-> (GET) http://{host}/{api-name}/{model}/{id}

To update a whole model:

-> (PUT) http://{host}/{api-name}/{model}/{id}

To update a partial content:

-> (PATCH) http://{host}/{api-name}/{model}/{id}

To delete a model:

-> (DELETE) http://{host}/{api-name}/{model}/{id}

