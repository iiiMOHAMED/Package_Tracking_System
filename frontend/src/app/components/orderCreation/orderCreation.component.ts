import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-orderCreation',
  templateUrl: './orderCreation.component.html',
})
export class orderCreationComponent {
    pickupLocation: string = '';
    dropOffLocation: string = '';
    packageDetails: string = '';
    deliveryTime: string = '';
  private apiUrl = 'http://localhost:8080/orders';

  constructor(private http: HttpClient) {}

  onSubmit() {
    const order = {
        pickupLocation: this.pickupLocation,
        dropOffLocation: this.dropOffLocation,
        packageDetails: this.packageDetails,
        deliveryTime: this.deliveryTime,
    };

    this.http.post(this.apiUrl, order).subscribe({
      next: () => alert('Order creation successfully!'),
      error: (err) => alert('Failed to create order: ' + err.error)
    });
  }
}
/*<!--type Order struct {
	OrderNumber     string `json:"orderNumber"`
	PickupLocation  string `json:"pickupLocation"`
	DropOffLocation string `json:"dropOffLocation"`
	PackageDetails  string `json:"packageDetails"`
	DeliveryTime    string `json:"deliveryTime"`
}
-->*/