<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Student extends Model
{
    protected $hidden = ["created_at","updated_at","deleted_at"];
    protected $guarded = ["id", "user_id"];
    protected $fillable = [
        'name',
        'user_id',
        'student_number',
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