import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class OrderService {
  private apiUrl = 'http://localhost:8080'; // Replace with your backend URL

  constructor(private http: HttpClient) {}

  getOrders(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/orders/retrieve`);
  }

  updateOrderStatus(orderId: number, newStatus: string): Observable<any> {
    return this.http.put(`${this.apiUrl}/orders/update/${orderId}`, { status: newStatus });
  }
 

  deleteOrder(orderId: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/orders/delete/${orderId}`);
  }

  assignCourierToOrder(orderId: number, courierId: number|null|{ courier_id: null }): Observable<any> {
    return this.http.put(`${this.apiUrl}/orders/assign/${orderId}`, { courier_id: courierId });
  }

  assignCourierToOrder2(orderId: number, courierId:{ courier_id: null }): Observable<any> {
    return this.http.put(`${this.apiUrl}/orders/assign/${orderId}`, { courier_id: courierId });
  }
   
   getUserOrders(userId: number): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/orders/users/${userId}`);
  }

  getOrderDetails(orderNumber: number): Observable<any> {
    return this.http.get(`${this.apiUrl}/orders/${orderNumber}`);
  }

  clearCourier(orderId: number, newStatus: string): Observable<any> {
    return this.http.put(`${this.apiUrl}/couriers/${orderId}`, { status: newStatus });
  }
  

}
