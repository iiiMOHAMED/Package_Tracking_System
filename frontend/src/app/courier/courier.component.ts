import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { OrderService } from '../order.service';
import { AuthService } from '../auth.service';
import { jwtDecode } from "jwt-decode";
import { CourierService } from '../courier.service';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-courier',
  templateUrl: './courier.component.html',
  styleUrls: ['./courier.component.css']
})
export class CourierComponent implements OnInit {
  pickupLocation: string = '';
  dropOffLocation: string = '';
  packageDetails: string = '';
  deliveryTime: string = '';
  private apiUrl = 'http://localhost:8080/orders';
  userId: number | null = null;
  ////////////////
  courierId: number= 0;
  orders: any[] = [];
  errorMessage: string = '';
  role:string='';
  constructor(private http: HttpClient, private courierService: CourierService,private orderService: OrderService, private authService: AuthService ) {
      const token = this.authService.getToken();
    if (token) {
      const decodedToken: any = jwtDecode(token);
      this.courierId = decodedToken.sub; 
      this.role=decodedToken.role;
      console.log("test1:"+this.courierId)
      console.log("test1:"+this.role)
    }
    else{ 
      console.log("test2:"+this.courierId)
      console.log("test2:"+this.role)
    }
  }
  isCourier(){
    if(this.role==="courier"){
      return true;
    }
    else{
      return false;
    }
  }

  ngOnInit(): void {
      this.load();
  }
  load(){
    // Get orders for the current user
    this.courierService.getCourierOrders(this.courierId).subscribe({
      next: (data) => {
        this.orders = data; // Successfully received orders
      },
      error: (error) => {
        this.errorMessage = 'There was an error retrieving orders.'; // Handle error
        console.error(error); // Log the error for debugging
      },
      complete: () => {
        console.log('Courier orders fetch complete'); // Optionally log completion
      }
    });
  }
  updateStatus(orderId: number, newStatus: string) {
    this.orderService.updateOrderStatus(orderId, newStatus).subscribe({
      next: response => {
        console.log('Status updated successfully', response);
        this.ngOnInit(); // Reload orders to reflect the updated status
        // Notify user of success
      },
      error: error => {
        console.error('Error updating order status:', error);
        alert('Error updating order status: ' + error.message); // Notify user of error
      }
    });
  }

  acceptOrder(orderId: number) {
    this.orderService.updateOrderStatus(orderId, "picked up").subscribe({
      next: response => {
        console.log('Order accepted successfully', response);
        this.ngOnInit(); // Reload orders to reflect the updated status
        // Notify user of success
      },
      error: error => {
        console.error('Error accepting order:', error);
        alert('Error accepting order: ' + error.message); // Notify user of error
      }
    });
  }

////////////////////////////
  acceptOrder2(orderId: number): void {
    this.courierService.acceptOrder(orderId).subscribe({
      next: () => {
        alert('Order accepted successfully!');
        this.ngOnInit(); // Refresh orders to update status
      },
      error: (error) => {
        console.error('Error accepting order:', error);
        console.log('ali:'+this.courierService.acceptOrder(orderId))
        alert('Error accepting order: ' + (error.error?.message || error.message));
      }
    });
  }
  clearCourier(orderId:number){
    this.orderService.clearCourier(orderId,'pending').subscribe({
      next: response => {
        console.log('Order declined successfully', response);
        this.ngOnInit(); // Reload orders to reflect the updated status
        alert('Order status updated successfully!'); // Notify user of success
      },
      error: error => {
        console.error('Error declined order:', error);
        alert('Error declined order: ' + error.message); // Notify user of error
      }
    })
  }

  declineOrder(orderId: number) {
    // Create a payload to update the courier_id to null
    const payload = { courier_id: null };  // courier_id is set to null
    this.orderService.assignCourierToOrder2(orderId,payload).subscribe({
      next: () => {
        this.ngOnInit(); // Refresh orders to update status
      },
      error: (error) => {
        console.error('Error declining order:', error);
        alert('Error declining order: ' + error.message);
      }
    });
  }
////////////////////////////
  revive(orderId: number,order:object){
    this.deleteOrder(orderId);
    this.reOrder(order);
  }

  deleteOrder(orderId: number) {
    this.orderService.deleteOrder(orderId).subscribe(response => {
      this.ngOnInit();
    });
  }
  reOrder(order:object) {
    
    this.http.post(this.apiUrl, order).subscribe({
      next: () => {
                  this.ngOnInit();
      },
      error: (err) => alert('Failed ' + err.error),
    });
  }


}
