import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { OrderService } from '../order.service';  // Adjust the path based on your project structure

@Component({
  selector: 'app-order-details',
  templateUrl: './order-details.component.html',
  styleUrls: ['./order-details.component.css']
})
export class OrderDetailsComponent implements OnInit {
  orderDetails: any = {};  // Store order details
  error: string = '';  // Store error message if any
  successMessage: string = '';  // Store success message after deleting order

  constructor(private route: ActivatedRoute, private orderService: OrderService, private router: Router) {}

  ngOnInit(): void {
    // Retrieve the orderNumber from the URL using snapshot
    const orderNumberStr = this.route.snapshot.paramMap.get('orderNumber');
    
    if (orderNumberStr) {
      const orderNumber = +orderNumberStr;
      if (orderNumber) {
        this.orderService.getOrderDetails(orderNumber).subscribe({
          next: (data) => {
            this.orderDetails = data;  // Set the order details
          },
          error: (err) => {
            this.error = 'Error fetching order details.';  // Error handling
          }
        });
      } else {
        this.error = 'Invalid order number.';  // If the orderNumber is invalid
      }
    } else {
      this.error = 'Order number not provided.';  // If no order number is found in URL
    }
  }

  deleteOrder(): void {
    if (this.orderDetails.status === 'pending') {
      // Call the service to delete the order
      this.orderService.deleteOrder(this.orderDetails.orderNumber).subscribe({
        next: () => {
          this.successMessage = 'Order canceled successfully.';
          this.orderDetails = {};  // Clear the order details after successful deletion
          this.router.navigate(['/user-orders']); // Redirect to the user's orders page after successful deletion
        },
        error: (err) => {
          this.error = 'Error canceling order.';
        }
      });
    } else {
      this.error = 'You can only cancel pending orders';
    }
  }
}
