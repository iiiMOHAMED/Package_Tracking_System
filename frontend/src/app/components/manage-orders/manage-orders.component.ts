import { Component, OnInit } from '@angular/core';
import { OrderService } from 'src/app/order.service';
@Component({
  selector: 'app-manage-orders',
  templateUrl: './manage-orders.component.html',
  styleUrls: ['./manage-orders.component.css']
})
export class ManageOrdersComponent implements OnInit {
  orders: any[] = [];
  orderNumber: number = 0; // Initialize with a default number
  courierId: number = 0; // Initialize with a default number
  constructor(private orderService: OrderService) {}

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
        alert('Order status updated successfully!'); // Notify user of success
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
        alert('Courier assigned successfully!'); // Notify user of success
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
