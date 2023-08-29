import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Observable, catchError, map, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  private authToken: string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbnRyZXByaXNlX2lkIjoxLCJleHAiOjE2OTAyOTM1NzksImlhdCI6MTY4NjY5MzU3OSwicm9sZV9ub20iOiJyb290IiwidXRpbGlzYXRldXJfaWQiOjF9.NtO41b-mgcYfuhTEZDtPZrphXtCeeB299xKv2UH9BHo";
  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      // Client-side error occurred
      console.error('An error occurred:', error.error.message);
    } else {
      // Server-side error occurred
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`
      );
    }
    // Return an observable with a user-facing error message
    return throwError('Failed to retrieve team data');
  }
  
  constructor(private http: HttpClient) {}
// Utilisateurs
  //GET
  getUtilisateur(): Observable<any> {
    const url = 'http://localhost:8080/api/app/utilisateur/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addUtilisateur(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/app/utilisateur/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  // DELETE
  deleteUtilisateur(id: any): Observable<any> {
    const url = `http://localhost:8080/api/app/utilisateur/${id}`;

    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editUtilisateur(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/app/utilisateur/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  
  

  
  
// Entreprise 
  //GET
  getEntreprise(): Observable<any> {
    const url = 'http://localhost:8080/api/app/entreprise/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
   //POST
   addEntreprise(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/app/entreprise/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteEntreprise(id: any): Observable<any> {
    const url = `http://localhost:8080/api/app/entreprise/${id}`;

    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editEntreprise(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/app/entreprise/${id}`;

    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }






//Audit
  //GET
  getAudit(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/audit/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addAudit(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/audit/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteAudit(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/audit/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editAudit(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/audit/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }


//Rapport
  
  //POST
  addRapport(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/rapport/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }




//Control List
  //GET
  getControlList(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/listecontrole/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addControlList(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/listecontrole/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteControlList(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/listecontrole/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //EDIT
  editControlList(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/listecontrole/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }





//Documents
  //GET
  getDocuments(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/revuedocument/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addDocuments(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/revuedocument/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteDocuments(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/revuedocument/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //EDIT
  editDocuments(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/revuedocument/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }



  
//Role
  //GET
  getRole(): Observable<any> {
    const url = 'http://localhost:8080/api/app/role/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }



//Requirement
  //GET
  getRequirement(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/exigence/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addRequirement(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/exigence/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteRequirement(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/exigence/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editRequirement(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/exigence/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }




//Meeting
  //GET
  getMeeting(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/reunion/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addMeeting(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/reunion/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteMeeting(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/reunion/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editMeeting(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/reunion/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }


//Tasks
  //GET
  getTask(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/tache/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addTask(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/tache/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteTask(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/tache/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editTask(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/tache/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }



//Perimeter
  //GET
  getPerimeter(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/perimetre/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addPerimeter(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/perimetre/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deletePerimeter(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/perimetre/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editPerimeter(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/perimetre/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }



//Objective
  //GET
  getObjective(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/objectif/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addObjective(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/objectif/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteObjective(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/objectif/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editObjective(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/objectif/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }




//Norm
  //GET
  getNorm(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/norme/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addNorm(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/norme/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteNorm(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/norme/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editNorm(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/norme/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }


//Criteria
  //GET
  getCriteria(): Observable<any> {
    const url = 'http://localhost:8080/api/v1/critere/all';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);

    return this.http.get(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  //POST
  addCriteria(data: any): Observable<any> {
    const url = 'http://localhost:8080/api/v1/critere/new';
    
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.post(url, data, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }

  // DELETE
  deleteCriteria(id: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/critere/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.delete(url, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
  editCriteria(id: any, formData: any): Observable<any> {
    const url = `http://localhost:8080/api/v1/critere/${id}`;
  
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + this.authToken);
  
    return this.http.put(url, formData, { headers: headers }).pipe(
      catchError(this.handleError)
    );
  }
}
