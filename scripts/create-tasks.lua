local counter = 0

request = function()
    counter = counter + 1
    headers = {}
    headers["Content-Type"] = "application/json"
    body = '{"storyId": 100, "title": "Task ' .. counter .. '"}'
    return wrk.format("POST", "/todos/api/v1/tasks", headers, body)
end
