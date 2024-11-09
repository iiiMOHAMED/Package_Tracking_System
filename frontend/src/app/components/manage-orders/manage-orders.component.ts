import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/auth.service';
import { OrderService } from 'src/app/order.service';
import { jwtDecode } from "jwt-decode";
@Component({
  selector: 'app-manage-orders',
  templateUrl: './manage-orders.component.html',
  styleUrls: ['./manage-orders.component.css']
})
export class ManageOrdersComponent implements OnInit {
  orders: any[] = [];
  orderNumber: number = 0; // Initialize with a default number
  courierId: number = 0; // Initialize with a default number
  role:string=" ";
  constructor(private orderService: OrderService, private authService: AuthService) {
    const token = this.authService.getToken();
    if (token) {
      const decodedToken: any = jwtDecode(token);
      this.role = decodedToken.role; 
      console.log("test1:"+this.role)
    }
    else{ 
      console.log("test2:"+this.role)
    }

  }

 isAdmin(){
  if(this.role==="admin"){
    return true;
  }
  else{
    return false;
  }
 }
 

  ngOnInit() {
    this.loadOrders();
  }

  loadOrders() {
    this.orderService.getOrders().subscribe(data => {
      this.orders = data;
    });
  }
 
  updateStatus(orderId: number, newStatus: string) {
    this.orderService.updateOrderStatus(orderId, newStatus).subscribe({
      next: response => {
        console.log('Status updated successfully', response);
        this.loadOrders(); // Reload orders to reflect the updated status
         // Notify user of success
      },
      error: error => {
        console.error('Error updating order status:', error);
        alert('Error updating order status: ' + error.message); // Notify user of error
      }
    });
  }

  assignCourier(orderId: number, courierId: number) {
    this.orderService.assignCourierToOrder(orderId, courierId).subscribe({
      next: response => {
        console.log('Courier assigned successfully', response);
        this.loadOrders(); // Reload orders to reflect the assigned courier
         // Notify user of success
      },
      error: error => {
        console.error('Error assigning courier:', error);
        alert('Error assigning courier: ' + error.message); // Notify user of error
      }
    });
  }
  
  

  deleteOrder(orderId: number) {
    this.orderService.deleteOrder(orderId).subscribe(response => {
      this.loadOrders();
    });
  }
}
