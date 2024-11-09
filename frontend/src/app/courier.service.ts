import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class CourierService {
  private apiUrl = 'http://localhost:8080'; // Replace with your backend URL

  constructor(private http: HttpClient) {}

  getCourierOrders(userId: number): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/orders/couriers/${userId}`);
  }

  // login(email: string, password: string): Observable<any> {
  //   const data = this.http.post<any>(`${this.apiUrl}/login`, {
  //     email,
  //     password,
  //   });
  //   return data;
  // }
  acceptOrder(orderId: number): Observable<any> {
    // Updates status to 'picked up' for the accepted order
    return this.http.put(`${this.apiUrl}/orders/accept/${orderId}`, {});
    
  }

  declineOrder(orderId: number): Observable<any> {
    // Sets courier_id to null and status to 'pending' for the declined order
    return this.http.put(`${this.apiUrl}/orders/decline/${orderId}`, {} );
  }
}
