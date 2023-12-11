<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Teacher extends Model
{
    protected $hidden = ["created_at","updated_at","deleted_at"];
    protected $guarded = ["id", "user_id"];
    protected $fillable = [
        'name',
        'user_id',
        'subject_id',
        'date_of_birth',
        'address',
        'no_handphone',
        'email',
        'employee_number',
    ];

    public function user()
    {
        return $this->belongsTo(User::class);
    }

    public function subject()
    {
        return $this->belongsTo(Subject::class);
    }
    
}
