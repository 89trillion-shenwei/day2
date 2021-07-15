from locust import HttpUser, between, task


class WebsiteUser(HttpUser):
    wait_time = between(5, 15)

    @task
    def index1(self):
        self.client.get("/Counter?string=323+232*132")
