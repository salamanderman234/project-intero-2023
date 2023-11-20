<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Student extends Model
{
    protected $fillable = [
        'name',
        'place_of_birth',
        'date_of_birth',
        'address',
        'no_handphone',
        'profile_pic',
    ];

    public function user()
    {
        return $this->belongsTo(User::class);
    }
}