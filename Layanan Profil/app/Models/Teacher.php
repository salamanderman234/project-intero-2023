<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Teacher extends Model
{
    protected $fillable = [
        'name',
        'date_of_birth',
        'address',
        'no_handphone',
        'email',
    ];

    public function user()
    {
        return $this->belongsTo(User::class);
    }
}
