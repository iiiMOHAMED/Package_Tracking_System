import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";

import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private apiUrl = 'http://localhost:8080'; 

  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<any> {
    const data = this.http.post<any>(`${this.apiUrl}/login`, {
      email,
      password,
    });
    return data;
  }

  storeToken(token: string): void {
    localStorage.setItem("token", token); // Store the token in local storage
  }

  getToken(): string | null {
    return localStorage.getItem("token"); // Retrieve the token
  }

  logout(): void {
    localStorage.removeItem("token"); // Remove the token
  }

  isLoggedIn(): boolean {
    return this.getToken() !== null; // Check if the user is logged in
  }

  // You can create an interceptor to add the token to each request if needed
  getHeaders(): HttpHeaders {
    return new HttpHeaders({
      Authorization: `Bearer ${this.getToken()}`, // Attach the token to the headers
    });
  }
}